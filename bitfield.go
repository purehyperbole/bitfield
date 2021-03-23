package bitfield

import (
	"errors"
)

// implements:
// https://stackoverflow.com/a/177092

// Bitfield represents a variable length bitfield
type Bitfield []uint8

// New creates a new bitfield of a given size in bytes
func New(size int) Bitfield {
	return make(Bitfield, size)
}

func (b Bitfield) Set(offset int) error {
	if offset < 0 || offset > (len(b)<<3)-1 {
		return errors.New("offset exceeds size of bitfield")
	}

	// get the correct byte to shift and then
	// construct a bitflag from the offset 8
	// bits long
	b[offset>>3] |= (1 << (offset & 0x7))

	return nil
}

func (b Bitfield) Has(offset int) bool {
	if offset < 0 || offset > (len(b)<<3)-1 {
		return false
	}

	return b[offset>>3]&(1<<(offset&0x7)) > 0
}

/*

int set_bit(char* bytes, unsigned long num_bytes, unsigned long offset)
{
  // make sure offset is valid
  if(offset < 0 || offset > (num_bytes<<3)-1) { return -1; }

  //set the right bit
  bytes[offset >> 3] |= (1 << (offset & 0x7));

  return 0; //success
}

//gets the n-th bit in |bytes|. num_bytes is the number of bytes in the array
// returns (-1) on error, 0 if bit is "off", positive number if "on"
int get_bit(char* bytes, unsigned long num_bytes, unsigned long offset)
{
  // make sure offset is valid
  if(offset < 0 || offset > (num_bytes<<3)-1) { return -1; }

  //get the right bit
  return (bytes[offset >> 3] & (1 << (offset & 0x7));
}

*/

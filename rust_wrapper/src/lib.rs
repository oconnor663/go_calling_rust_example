use core::ffi::c_char;
use core::ptr;
use core::slice;

#[no_mangle]
pub extern "C" fn blake3_wrapper(input: *const c_char, input_len: usize, output: *mut c_char) {
    let input_slice = unsafe { slice::from_raw_parts(input as *const u8, input_len) };
    let hash = blake3::hash(input_slice);
    unsafe {
        ptr::write(output as *mut [u8; 32], *hash.as_bytes());
    }
}

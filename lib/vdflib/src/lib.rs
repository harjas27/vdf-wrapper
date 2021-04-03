extern crate vdf;
use vdf::{VDFParams, WesolowskiVDFParams, VDF};

extern crate libc;


#[no_mangle]
pub extern "C" fn compute(
    array_pointer: *const u8,
    size: libc::size_t,
    ans_ptr: *mut u8,
    ans_size: libc::size_t,
    difficulty: u64,
    num_bits: u16
){
    let pietrzak_vdf = WesolowskiVDFParams(num_bits).new();
    let sol: &[u8] = unsafe {
            std::slice::from_raw_parts(array_pointer as *const u8, size as usize)
        };
    let ans: &mut [u8] = unsafe {
        std::slice::from_raw_parts_mut(ans_ptr as *mut u8, ans_size as usize)
    };
    copy_slice(ans, &pietrzak_vdf.solve(sol, difficulty).unwrap()[..]);
}

#[no_mangle]
pub extern "C" fn verify(
    array_pointer: *const u8,
    size: libc::size_t,
    ans_ptr: *const u8,
    ans_size: libc::size_t,
    difficulty: u64,
    num_bits: u16
) -> i32 {
    let pietrzak_vdf = WesolowskiVDFParams(num_bits).new();
    let sol: &[u8] = unsafe {
        std::slice::from_raw_parts(array_pointer as *const u8, size as usize)
    };
    let ans: &[u8] = unsafe {
        std::slice::from_raw_parts(ans_ptr as *const u8, ans_size as usize)
    };
    return if pietrzak_vdf.verify(sol, difficulty, ans).is_ok() {
        1
    } else {
        0
    }
}


fn copy_slice(dst: &mut [u8], src: &[u8]) -> usize {
    let mut c = 0;
    for (d, s) in dst.iter_mut().zip(src.iter()) {
        *d = *s;
        c += 1;
    }
    c
}
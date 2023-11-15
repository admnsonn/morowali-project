<?php

namespace App\Http\Controllers;

use App\Http\Controllers\Controller;
use Illuminate\Support\Facades\DB;
use Illuminate\Http\Request;
use App\Models\Pengguna;

class PenggunaController extends Controller
{
    public function index()
    {
        $pengguna = Pengguna::with('role')->get();
        return response()->json($pengguna, 200);
    }

    public function login(Request $request)
    {
         $nik = $request->input('nik');
         $password = $request->input('password');
 
         $pengguna = DB::table('pengguna')
             ->where('nik', $nik)
             ->where('password', $password)
             ->first();
 
         if ($pengguna) {
            
            $penggunaArray = array_map('strval', (array) $pengguna);

             return response()->json([
                'status' => true,
                'message'=> 'Berhasil Login',
                'data' => $penggunaArray
            ], 
            200);

         } else {
            return response()->json([
                'status' => false,
                'message'=> 'Login Gagal',
                'data' => []
            ], 
            200);
         }
    }
}

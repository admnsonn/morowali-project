<?php


namespace App\Models;

use Illuminate\Database\Eloquent\Model;

class Pengguna extends Model
{
    protected $table = 'pengguna';

    // Definisi relasi dengan tabel Role
    public function role()
    {
        return $this->belongsTo(Role::class, 'id_role');
    }
}
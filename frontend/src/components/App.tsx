import React, { useState, useEffect } from 'react';
import { Auth } from './Auth';
import { SandwichList } from './SandwichList';
import { OrderList } from './OrderList';
import { LogOut, User as UserIcon, Sandwich as SandwichIcon, ClipboardList } from 'lucide-react';

export const App: React.FC = () => {
  const [user, setUser] = useState<{ email: string; role: string } | null>(null);
  const [view, setView] = useState<'menu' | 'orders'>('menu');

  useEffect(() => {
    const savedUser = localStorage.getItem('user');
    if (savedUser) {
      setUser(JSON.parse(savedUser));
    }
  }, []);

  const handleLogout = () => {
    localStorage.removeItem('token');
    localStorage.removeItem('user');
    setUser(null);
  };

  if (!user) {
    return (
      <div style={{ marginTop: '5rem' }}>
        <h1 style={{ textAlign: 'center', marginBottom: '2rem' }}>🥪 PaninApp</h1>
        <Auth onLogin={setUser} />
      </div>
    );
  }

  return (
    <div style={{ display: 'flex', flexDirection: 'column', gap: '2rem' }}>
      <header style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', padding: '1rem', background: '#23262d', borderRadius: '8px' }}>
        <div style={{ display: 'flex', alignItems: 'center', gap: '0.5rem' }}>
          <UserIcon size={20} />
          <span>{user.email} ({user.role})</span>
        </div>
        <button onClick={handleLogout} style={{ display: 'flex', alignItems: 'center', gap: '0.5rem', background: 'none', border: '1px solid #444', color: 'white', padding: '0.5rem', borderRadius: '4px', cursor: 'pointer' }}>
          <LogOut size={16} /> Logout
        </button>
      </header>

      <nav style={{ display: 'flex', gap: '1rem', borderBottom: '1px solid #333', paddingBottom: '1rem' }}>
        <button
          onClick={() => setView('menu')}
          style={{
            flex: 1, padding: '0.8rem', display: 'flex', alignItems: 'center', justifyContent: 'center', gap: '0.5rem',
            background: view === 'menu' ? 'var(--accent)' : 'none',
            border: 'none', color: 'white', borderRadius: '4px', cursor: 'pointer', fontWeight: 'bold'
          }}
        >
          <SandwichIcon size={18} /> Menu
        </button>
        <button
          onClick={() => setView('orders')}
          style={{
            flex: 1, padding: '0.8rem', display: 'flex', alignItems: 'center', justifyContent: 'center', gap: '0.5rem',
            background: view === 'orders' ? 'var(--accent)' : 'none',
            border: 'none', color: 'white', borderRadius: '4px', cursor: 'pointer', fontWeight: 'bold'
          }}
        >
          <ClipboardList size={18} /> Orders
        </button>
      </nav>

      {view === 'menu' ? <SandwichList role={user.role} /> : <OrderList />}
    </div>
  );
};

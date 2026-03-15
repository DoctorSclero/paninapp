import React, { useEffect, useState } from 'react';
import api from '../api/client';
import { ShoppingCart, Plus, Minus } from 'lucide-react';

interface Sandwich {
  ID: number;
  name: string;
  description: string;
  price: number;
}

interface CartItem {
  sandwich: Sandwich;
  quantity: number;
}

export const SandwichList: React.FC<{ role: string }> = ({ role }) => {
  const [sandwiches, setSandwiches] = useState<Sandwich[]>([]);
  const [cart, setCart] = useState<CartItem[]>([]);
  const [message, setMessage] = useState('');
  const [newSandwich, setNewSandwich] = useState({ name: '', description: '', price: 0 });

  useEffect(() => {
    fetchSandwiches();
  }, []);

  const fetchSandwiches = async () => {
    try {
      const response = await api.get('/sandwiches');
      setSandwiches(response.data.sandwiches || []);
    } catch (err) {
      console.error('Failed to fetch sandwiches', err);
    }
  };

  const addToCart = (sw: Sandwich) => {
    setCart((prev) => {
      const existing = prev.find((item) => item.sandwich.ID === sw.ID);
      if (existing) {
        return prev.map((item) =>
          item.sandwich.ID === sw.ID ? { ...item, quantity: item.quantity + 1 } : item
        );
      }
      return [...prev, { sandwich: sw, quantity: 1 }];
    });
  };

  const removeFromCart = (swId: number) => {
    setCart((prev) => {
      const existing = prev.find((item) => item.sandwich.ID === swId);
      if (existing && existing.quantity > 1) {
        return prev.map((item) =>
          item.sandwich.ID === swId ? { ...item, quantity: item.quantity - 1 } : item
        );
      }
      return prev.filter((item) => item.sandwich.ID !== swId);
    });
  };

  const placeOrder = async () => {
    try {
      const items = cart.map((item) => ({
        sandwich_id: item.sandwich.ID,
        quantity: item.quantity,
      }));
      await api.post('/orders', { items });
      setCart([]);
      setMessage('Order placed successfully!');
    } catch (err) {
      setMessage('Failed to place order.');
    }
  };

  const createSandwich = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await api.post('/sandwiches', newSandwich);
      setNewSandwich({ name: '', description: '', price: 0 });
      fetchSandwiches();
      setMessage('Sandwich created!');
    } catch (err) {
      setMessage('Failed to create sandwich.');
    }
  };

  const total = cart.reduce((sum, item) => sum + item.sandwich.price * item.quantity, 0);

  return (
    <div style={{ display: 'flex', flexDirection: 'column', gap: '2rem' }}>
      {role === 'manager' && (
        <div style={{ background: '#23262d', padding: '1.5rem', borderRadius: '8px' }}>
          <h3>Add New Sandwich</h3>
          <form onSubmit={createSandwich} style={{ display: 'flex', gap: '1rem', flexWrap: 'wrap' }}>
            <input
              type="text"
              placeholder="Name"
              value={newSandwich.name}
              onChange={(e) => setNewSandwich({ ...newSandwich, name: e.target.value })}
              required
              style={{ padding: '0.5rem' }}
            />
            <input
              type="text"
              placeholder="Description"
              value={newSandwich.description}
              onChange={(e) => setNewSandwich({ ...newSandwich, description: e.target.value })}
              style={{ padding: '0.5rem' }}
            />
            <input
              type="number"
              placeholder="Price"
              step="0.01"
              value={newSandwich.price || ''}
              onChange={(e) => setNewSandwich({ ...newSandwich, price: parseFloat(e.target.value) })}
              required
              style={{ padding: '0.5rem' }}
            />
            <button type="submit" style={{ padding: '0.5rem 1rem', background: 'green', color: 'white', border: 'none', borderRadius: '4px', cursor: 'pointer' }}>
              Create
            </button>
          </form>
        </div>
      )}

      <div style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fill, minmax(200px, 1fr))', gap: '1rem' }}>
        {sandwiches.map((sw) => (
          <div key={sw.ID} style={{ background: '#23262d', padding: '1rem', borderRadius: '8px', border: '1px solid #444' }}>
            <h4>{sw.name}</h4>
            <p style={{ fontSize: '0.9rem', color: '#ccc' }}>{sw.description}</p>
            <p style={{ fontWeight: 'bold' }}>€{sw.price.toFixed(2)}</p>
            <button
              onClick={() => addToCart(sw)}
              style={{ width: '100%', padding: '0.5rem', background: '#3b82f6', color: 'white', border: 'none', borderRadius: '4px', cursor: 'pointer', display: 'flex', alignItems: 'center', justifyContent: 'center', gap: '0.5rem' }}
            >
              <Plus size={16} /> Add to Order
            </button>
          </div>
        ))}
      </div>

      {cart.length > 0 && (
        <div style={{ position: 'sticky', bottom: '1rem', background: '#23262d', padding: '1.5rem', borderRadius: '8px', border: '2px solid var(--accent)', boxShadow: '0 -4px 10px rgba(0,0,0,0.5)' }}>
          <h3 style={{ display: 'flex', alignItems: 'center', gap: '0.5rem' }}>
            <ShoppingCart /> Your Order
          </h3>
          <ul style={{ listStyle: 'none', padding: 0 }}>
            {cart.map((item) => (
              <li key={item.sandwich.ID} style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '0.5rem' }}>
                <span>{item.sandwich.name} x{item.quantity}</span>
                <div style={{ display: 'flex', gap: '0.5rem', alignItems: 'center' }}>
                  <span>€{(item.sandwich.price * item.quantity).toFixed(2)}</span>
                  <button onClick={() => removeFromCart(item.sandwich.ID)} style={{ background: 'red', border: 'none', color: 'white', borderRadius: '4px', cursor: 'pointer', padding: '0.2rem' }}>
                    <Minus size={14} />
                  </button>
                </div>
              </li>
            ))}
          </ul>
          <div style={{ borderTop: '1px solid #444', paddingTop: '1rem', marginTop: '1rem', display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
            <strong>Total: €{total.toFixed(2)}</strong>
            <button onClick={placeOrder} style={{ padding: '0.7rem 1.5rem', background: 'var(--accent)', color: 'white', border: 'none', borderRadius: '4px', cursor: 'pointer', fontWeight: 'bold' }}>
              Confirm Order
            </button>
          </div>
        </div>
      )}

      {message && <p style={{ color: 'lightblue', textAlign: 'center' }}>{message}</p>}
    </div>
  );
};

import React, { useEffect, useState } from 'react';
import api from '../api/client';
import { Clock, CheckCircle, Package } from 'lucide-react';

interface OrderItem {
  ID: number;
  sandwich: { name: string };
  quantity: number;
  price: number;
}

interface Order {
  ID: number;
  status: string;
  total: number;
  CreatedAt: string;
  items: OrderItem[];
}

export const OrderList: React.FC = () => {
  const [orders, setOrders] = useState<Order[]>([]);

  useEffect(() => {
    const interval = setInterval(fetchOrders, 5000);
    fetchOrders();
    return () => clearInterval(interval);
  }, []);

  const fetchOrders = async () => {
    try {
      const response = await api.get('/orders');
      setOrders(response.data.orders || []);
    } catch (err) {
      console.error('Failed to fetch orders', err);
    }
  };

  const getStatusIcon = (status: string) => {
    switch (status) {
      case 'pending': return <Clock size={16} color="orange" />;
      case 'confirmed': return <Package size={16} color="lightblue" />;
      case 'completed': return <CheckCircle size={16} color="green" />;
      default: return null;
    }
  };

  return (
    <div style={{ display: 'flex', flexDirection: 'column', gap: '1rem' }}>
      <h3>Orders</h3>
      {orders.length === 0 && <p>No active orders.</p>}
      {orders.map((order) => (
        <div key={order.ID} style={{ background: '#23262d', padding: '1rem', borderRadius: '8px', borderLeft: `4px solid ${order.status === 'completed' ? 'green' : 'orange'}` }}>
          <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
            <span style={{ display: 'flex', alignItems: 'center', gap: '0.5rem', textTransform: 'uppercase', fontSize: '0.8rem', fontWeight: 'bold' }}>
              {getStatusIcon(order.status)} {order.status}
            </span>
            <span style={{ fontSize: '0.8rem', color: '#888' }}>
              {new Date(order.CreatedAt).toLocaleString()}
            </span>
          </div>
          <div style={{ marginTop: '0.5rem' }}>
            {order.items?.map((item) => (
              <div key={item.ID} style={{ fontSize: '0.9rem', color: '#ccc' }}>
                {item.sandwich?.name} x{item.quantity}
              </div>
            ))}
          </div>
          <div style={{ marginTop: '0.5rem', textAlign: 'right', fontWeight: 'bold' }}>
            Total: €{order.total.toFixed(2)}
          </div>
        </div>
      ))}
    </div>
  );
};

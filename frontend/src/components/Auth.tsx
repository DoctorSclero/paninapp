import React, { useState } from 'react';
import api from '../api/client';

interface AuthProps {
  onLogin: (user: { email: string; role: string }) => void;
}

export const Auth: React.FC<AuthProps> = ({ onLogin }) => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [role, setRole] = useState('consumer');
  const [isRegistering, setIsRegistering] = useState(false);
  const [error, setError] = useState('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');
    const endpoint = isRegistering ? '/users/register' : '/users/login';
    const payload = isRegistering ? { email, password, role } : { email, password };

    try {
      const response = await api.post(endpoint, payload);
      const { token, user: userData, role: userRole } = response.data;
      
      localStorage.setItem('token', token);
      
      // If registering, the backend returns the full user object in 'user'
      // If logging in, the backend returns the email in 'user' and the role in 'role'
      const emailValue = typeof userData === 'string' ? userData : userData.Email;
      const roleValue = userRole || userData.role || (isRegistering ? role : 'unknown');

      const finalUser = {
        email: emailValue,
        role: roleValue
      };

      localStorage.setItem('user', JSON.stringify(finalUser));
      onLogin(finalUser);
    } catch (err: any) {
      setError(err.response?.data?.error || 'Authentication failed');
    }
  };

  return (
    <div style={{ maxWidth: '400px', margin: 'auto', padding: '2rem', background: '#23262d', borderRadius: '8px' }}>
      <h2>{isRegistering ? 'Create Account' : 'Login'}</h2>
      <form onSubmit={handleSubmit} style={{ display: 'flex', flexDirection: 'column', gap: '1rem' }}>
        <input
          type="email"
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          required
          style={{ padding: '0.5rem' }}
        />
        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
          style={{ padding: '0.5rem' }}
        />
        {isRegistering && (
          <select value={role} onChange={(e) => setRole(e.target.value)} style={{ padding: '0.5rem' }}>
            <option value="consumer">Consumer</option>
            <option value="manager">Manager</option>
          </select>
        )}
        <button type="submit" style={{ padding: '0.5rem', background: 'var(--accent)', color: 'white', border: 'none', borderRadius: '4px', cursor: 'pointer' }}>
          {isRegistering ? 'Register' : 'Login'}
        </button>
      </form>
      {error && <p style={{ color: 'red', marginTop: '1rem' }}>{error}</p>}
      <p style={{ marginTop: '1rem', textAlign: 'center' }}>
        <button onClick={() => setIsRegistering(!isRegistering)} style={{ background: 'none', border: 'none', color: 'lightblue', cursor: 'pointer', textDecoration: 'underline' }}>
          {isRegistering ? 'Already have an account? Login' : "Don't have an account? Register"}
        </button>
      </p>
    </div>
  );
};

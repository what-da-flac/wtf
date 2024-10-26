import React from 'react';
import ConfettiExplosion from 'react-confetti-explosion';

type params = {
  isExploding: boolean;
};
export default function SuccessAnimation({ isExploding }: params) {
  return (
    <React.Fragment>
      {isExploding && (
        <div
          style={{
            position: 'absolute',
            top: '20%',
            left: '55%',
          }}
        >
          <ConfettiExplosion
            colors={['#FFC700', '#FF0000', '#2E3191', '#1971c2']}
            particleCount={400}
            duration={5000}
          />
        </div>
      )}
    </React.Fragment>
  );
}

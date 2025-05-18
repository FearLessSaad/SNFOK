import { Skeleton } from '@/components/ui/skeleton'
import React from 'react'

function getRandomWidth() {
  // Returns a random width class between 40% and 90%
  const widths = ['w-2/5', 'w-1/2', 'w-3/5', 'w-2/3', 'w-3/4', 'w-4/5', 'w-5/6', 'w-11/12'];
  return widths[Math.floor(Math.random() * widths.length)];
}

function LoadingSkeleton() {
  const cardCount = 4;
  const cards = Array.from({ length: cardCount });

  return (
    <div className="grid grid-cols-1 md:grid-cols-4 gap-2 w-full">
      {cards.map((_, cardIdx) => (
        <div
          key={cardIdx}
          className="bg-muted rounded-xl shadow-sm p-5 flex flex-col gap-4 min-h-[180px] border border-muted-foreground/10"
        >
          {[...Array(4)].map((__, idx) => (
            <Skeleton
              key={idx}
              className={`h-7 ${getRandomWidth()} mb-2`}
            />
          ))}
        </div>
      ))}
    </div>
  );
}

export default LoadingSkeleton

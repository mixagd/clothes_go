import { TestBed, inject } from '@angular/core/testing';

import { ClothesService } from './clothes.service';

describe('ClothesService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [ClothesService]
    });
  });

  it('should be created', inject([ClothesService], (service: ClothesService) => {
    expect(service).toBeTruthy();
  }));
});

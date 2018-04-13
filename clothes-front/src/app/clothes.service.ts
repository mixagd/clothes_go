import { Injectable } from '@angular/core';
import {Cloth} from './model';
import {Observable} from 'rxjs/Observable';
import {HttpClient} from '@angular/common/http';

@Injectable()
export class ClothesService {

  constructor(private http: HttpClient) {}

  getAllClothes(): Observable<Cloth[]> {
    return this.http.get<Cloth[]>('/api/clothes');
  }

}

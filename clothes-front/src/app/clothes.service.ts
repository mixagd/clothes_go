import { Injectable } from '@angular/core';
import {Cloth} from './model';
import {Observable} from 'rxjs/Observable';
import {HttpClient} from '@angular/common/http';
import {ClothFormData} from './cloth-form/cloth-form.component';

@Injectable()
export class ClothesService {

  constructor(private http: HttpClient) {}

  getAllClothes(): Observable<Cloth[]> {
    return this.http.get<Cloth[]>('/api/clothes');
  }

  createNewCloth(cl: ClothFormData): Observable<Cloth> {
    return this.http.post<Cloth>('/api/clothes/new', {type: cl.type, colour: cl.colour, fit: cl.fit, owner: cl.owner});
  }

  getCloth(id: string): Observable<Cloth> {
    return this.http.get<Cloth>(`/api/clothes/${id}`);
  }

  updateCloth(id: string, cl: ClothFormData): Observable<Cloth> {
    return this.http.put<Cloth>(`/api/clothes/${id}`, {type: cl.type, colour: cl.colour, fit: cl.fit, owner: cl.owner});
  }

}

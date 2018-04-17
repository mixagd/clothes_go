import { Component, OnInit } from '@angular/core';
import {Observable} from 'rxjs/Observable';
import {Cloth} from '../model';
import {ClothesService} from '../clothes.service';

@Component({
  selector: 'app-clothes',
  templateUrl: './clothes.component.html',
  styleUrls: ['./clothes.component.css']
})
export class ClothesComponent implements OnInit {

  clothes$: Observable<Cloth[]>;

  constructor(private clothesService: ClothesService) {}

  ngOnInit() {
    this.clothes$ = this.clothesService.getAllClothes();
    console.log('lalala');
  }

}

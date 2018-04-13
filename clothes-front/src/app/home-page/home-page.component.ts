import { Component, OnInit } from '@angular/core';
import {AppComponent} from '../app.component';
import {Observable} from 'rxjs/Observable';
import {Cloth} from '../model';
import {ClothesService} from '../clothes.service';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css']
})
export class HomePageComponent implements OnInit {

  clothes$: Observable<Cloth[]>;

  constructor(private clothesService: ClothesService) {}

  ngOnInit() {
    this.clothes$ = this.clothesService.getAllClothes();
  }

}

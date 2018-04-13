import {Component, OnInit} from '@angular/core';
import {ClothesService} from './clothes.service';
import {Observable} from 'rxjs/Observable';
import {Cloth} from './model';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  title = 'Clothes';

  clothes$: Observable<Cloth[]>;

  constructor(private clothesService: ClothesService) {}

  ngOnInit() {
    this.clothes$ = this.clothesService.getAllClothes();
  }
}

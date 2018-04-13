import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';


import { AppComponent } from './app.component';
import { HomePageComponent } from './home-page/home-page.component';
import {ClothesService} from './clothes.service';
import {HttpClientModule} from '@angular/common/http';
import { ClothItemComponent } from './cloth-item/cloth-item.component';


@NgModule({
  declarations: [
    AppComponent,
    HomePageComponent,
    ClothItemComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule
  ],
  providers: [ClothesService],
  bootstrap: [AppComponent]
})
export class AppModule { }

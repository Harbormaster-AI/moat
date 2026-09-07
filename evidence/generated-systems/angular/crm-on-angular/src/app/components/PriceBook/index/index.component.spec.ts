
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexPriceBookComponent } from './index.component';
import { PriceBookService } from '../../../services/PriceBook.service';

describe('IndexPriceBookComponent', () => {
  let component: IndexPriceBookComponent;
  let fixture: ComponentFixture<IndexPriceBookComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexPriceBookComponent
      ],
      providers: [
        PriceBookService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexPriceBookComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
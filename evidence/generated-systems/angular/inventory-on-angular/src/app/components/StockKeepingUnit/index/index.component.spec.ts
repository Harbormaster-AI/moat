
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexStockKeepingUnitComponent } from './index.component';
import { StockKeepingUnitService } from '../../../services/StockKeepingUnit.service';

describe('IndexStockKeepingUnitComponent', () => {
  let component: IndexStockKeepingUnitComponent;
  let fixture: ComponentFixture<IndexStockKeepingUnitComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexStockKeepingUnitComponent
      ],
      providers: [
        StockKeepingUnitService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexStockKeepingUnitComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
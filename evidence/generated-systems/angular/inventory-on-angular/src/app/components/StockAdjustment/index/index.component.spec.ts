
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexStockAdjustmentComponent } from './index.component';
import { StockAdjustmentService } from '../../../services/StockAdjustment.service';

describe('IndexStockAdjustmentComponent', () => {
  let component: IndexStockAdjustmentComponent;
  let fixture: ComponentFixture<IndexStockAdjustmentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexStockAdjustmentComponent
      ],
      providers: [
        StockAdjustmentService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexStockAdjustmentComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
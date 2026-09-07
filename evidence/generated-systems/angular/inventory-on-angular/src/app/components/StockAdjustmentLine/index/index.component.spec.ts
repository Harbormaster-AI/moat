
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexStockAdjustmentLineComponent } from './index.component';
import { StockAdjustmentLineService } from '../../../services/StockAdjustmentLine.service';

describe('IndexStockAdjustmentLineComponent', () => {
  let component: IndexStockAdjustmentLineComponent;
  let fixture: ComponentFixture<IndexStockAdjustmentLineComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexStockAdjustmentLineComponent
      ],
      providers: [
        StockAdjustmentLineService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexStockAdjustmentLineComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
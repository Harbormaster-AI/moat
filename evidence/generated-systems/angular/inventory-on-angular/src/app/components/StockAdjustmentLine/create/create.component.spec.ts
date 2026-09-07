
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateStockAdjustmentLineComponent } from './create.component';
import { StockAdjustmentLineService } from '../../../services/StockAdjustmentLine.service';
import { Router } from '@angular/router';

describe('CreateStockAdjustmentLineComponent', () => {
  let component: CreateStockAdjustmentLineComponent;
  let fixture: ComponentFixture<CreateStockAdjustmentLineComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateStockAdjustmentLineComponent
      ],
      providers: [
        StockAdjustmentLineService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateStockAdjustmentLineComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
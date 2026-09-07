
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateStockAdjustmentComponent } from './create.component';
import { StockAdjustmentService } from '../../../services/StockAdjustment.service';
import { Router } from '@angular/router';

describe('CreateStockAdjustmentComponent', () => {
  let component: CreateStockAdjustmentComponent;
  let fixture: ComponentFixture<CreateStockAdjustmentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateStockAdjustmentComponent
      ],
      providers: [
        StockAdjustmentService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateStockAdjustmentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
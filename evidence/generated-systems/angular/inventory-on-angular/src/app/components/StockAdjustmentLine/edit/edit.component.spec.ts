
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditStockAdjustmentLineComponent } from './edit.component';
import { StockAdjustmentLineService } from '../../../services/StockAdjustmentLine.service';

describe('EditStockAdjustmentLineComponent', () => {
  let component: EditStockAdjustmentLineComponent;
  let fixture: ComponentFixture<EditStockAdjustmentLineComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditStockAdjustmentLineComponent
      ],
      providers: [
        StockAdjustmentLineService,
        {
          provide: ActivatedRoute,
          useValue: {
            params: of({ id: '1' })
          }
        },
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(EditStockAdjustmentLineComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
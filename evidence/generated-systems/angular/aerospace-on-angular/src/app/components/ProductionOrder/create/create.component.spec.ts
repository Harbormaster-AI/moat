
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateProductionOrderComponent } from './create.component';
import { ProductionOrderService } from '../../../services/ProductionOrder.service';
import { Router } from '@angular/router';

describe('CreateProductionOrderComponent', () => {
  let component: CreateProductionOrderComponent;
  let fixture: ComponentFixture<CreateProductionOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateProductionOrderComponent
      ],
      providers: [
        ProductionOrderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateProductionOrderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
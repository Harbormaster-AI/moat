
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreatePaymentOrderComponent } from './create.component';
import { PaymentOrderService } from '../../../services/PaymentOrder.service';
import { Router } from '@angular/router';

describe('CreatePaymentOrderComponent', () => {
  let component: CreatePaymentOrderComponent;
  let fixture: ComponentFixture<CreatePaymentOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreatePaymentOrderComponent
      ],
      providers: [
        PaymentOrderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreatePaymentOrderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
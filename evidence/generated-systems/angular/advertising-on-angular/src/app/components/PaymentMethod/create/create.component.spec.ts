
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreatePaymentMethodComponent } from './create.component';
import { PaymentMethodService } from '../../../services/PaymentMethod.service';
import { Router } from '@angular/router';

describe('CreatePaymentMethodComponent', () => {
  let component: CreatePaymentMethodComponent;
  let fixture: ComponentFixture<CreatePaymentMethodComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreatePaymentMethodComponent
      ],
      providers: [
        PaymentMethodService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreatePaymentMethodComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
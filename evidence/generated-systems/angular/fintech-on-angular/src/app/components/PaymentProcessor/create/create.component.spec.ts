
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreatePaymentProcessorComponent } from './create.component';
import { PaymentProcessorService } from '../../../services/PaymentProcessor.service';
import { Router } from '@angular/router';

describe('CreatePaymentProcessorComponent', () => {
  let component: CreatePaymentProcessorComponent;
  let fixture: ComponentFixture<CreatePaymentProcessorComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreatePaymentProcessorComponent
      ],
      providers: [
        PaymentProcessorService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreatePaymentProcessorComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
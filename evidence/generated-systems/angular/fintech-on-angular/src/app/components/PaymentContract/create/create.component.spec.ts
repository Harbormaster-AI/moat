
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreatePaymentContractComponent } from './create.component';
import { PaymentContractService } from '../../../services/PaymentContract.service';
import { Router } from '@angular/router';

describe('CreatePaymentContractComponent', () => {
  let component: CreatePaymentContractComponent;
  let fixture: ComponentFixture<CreatePaymentContractComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreatePaymentContractComponent
      ],
      providers: [
        PaymentContractService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreatePaymentContractComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
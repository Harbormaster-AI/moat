
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexPaymentContractComponent } from './index.component';
import { PaymentContractService } from '../../../services/PaymentContract.service';

describe('IndexPaymentContractComponent', () => {
  let component: IndexPaymentContractComponent;
  let fixture: ComponentFixture<IndexPaymentContractComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexPaymentContractComponent
      ],
      providers: [
        PaymentContractService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexPaymentContractComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
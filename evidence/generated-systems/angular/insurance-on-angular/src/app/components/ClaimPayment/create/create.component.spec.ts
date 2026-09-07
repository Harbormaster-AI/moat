
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateClaimPaymentComponent } from './create.component';
import { ClaimPaymentService } from '../../../services/ClaimPayment.service';
import { Router } from '@angular/router';

describe('CreateClaimPaymentComponent', () => {
  let component: CreateClaimPaymentComponent;
  let fixture: ComponentFixture<CreateClaimPaymentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateClaimPaymentComponent
      ],
      providers: [
        ClaimPaymentService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateClaimPaymentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
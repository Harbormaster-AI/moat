
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditPurchaseAgreementComponent } from './edit.component';
import { PurchaseAgreementService } from '../../../services/PurchaseAgreement.service';

describe('EditPurchaseAgreementComponent', () => {
  let component: EditPurchaseAgreementComponent;
  let fixture: ComponentFixture<EditPurchaseAgreementComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditPurchaseAgreementComponent
      ],
      providers: [
        PurchaseAgreementService,
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

    fixture = TestBed.createComponent(EditPurchaseAgreementComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
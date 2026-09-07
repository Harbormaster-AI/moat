
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateReinsuranceAgreementComponent } from './create.component';
import { ReinsuranceAgreementService } from '../../../services/ReinsuranceAgreement.service';
import { Router } from '@angular/router';

describe('CreateReinsuranceAgreementComponent', () => {
  let component: CreateReinsuranceAgreementComponent;
  let fixture: ComponentFixture<CreateReinsuranceAgreementComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateReinsuranceAgreementComponent
      ],
      providers: [
        ReinsuranceAgreementService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateReinsuranceAgreementComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
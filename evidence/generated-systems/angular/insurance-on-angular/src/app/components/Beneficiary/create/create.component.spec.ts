
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateBeneficiaryComponent } from './create.component';
import { BeneficiaryService } from '../../../services/Beneficiary.service';
import { Router } from '@angular/router';

describe('CreateBeneficiaryComponent', () => {
  let component: CreateBeneficiaryComponent;
  let fixture: ComponentFixture<CreateBeneficiaryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateBeneficiaryComponent
      ],
      providers: [
        BeneficiaryService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateBeneficiaryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
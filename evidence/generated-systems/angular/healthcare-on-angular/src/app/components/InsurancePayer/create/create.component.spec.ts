
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateInsurancePayerComponent } from './create.component';
import { InsurancePayerService } from '../../../services/InsurancePayer.service';
import { Router } from '@angular/router';

describe('CreateInsurancePayerComponent', () => {
  let component: CreateInsurancePayerComponent;
  let fixture: ComponentFixture<CreateInsurancePayerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateInsurancePayerComponent
      ],
      providers: [
        InsurancePayerService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateInsurancePayerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
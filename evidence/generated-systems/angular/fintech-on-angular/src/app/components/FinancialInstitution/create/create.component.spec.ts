
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateFinancialInstitutionComponent } from './create.component';
import { FinancialInstitutionService } from '../../../services/FinancialInstitution.service';
import { Router } from '@angular/router';

describe('CreateFinancialInstitutionComponent', () => {
  let component: CreateFinancialInstitutionComponent;
  let fixture: ComponentFixture<CreateFinancialInstitutionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateFinancialInstitutionComponent
      ],
      providers: [
        FinancialInstitutionService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateFinancialInstitutionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
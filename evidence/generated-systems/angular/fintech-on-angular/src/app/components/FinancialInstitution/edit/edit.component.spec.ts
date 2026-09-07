
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditFinancialInstitutionComponent } from './edit.component';
import { FinancialInstitutionService } from '../../../services/FinancialInstitution.service';

describe('EditFinancialInstitutionComponent', () => {
  let component: EditFinancialInstitutionComponent;
  let fixture: ComponentFixture<EditFinancialInstitutionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditFinancialInstitutionComponent
      ],
      providers: [
        FinancialInstitutionService,
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

    fixture = TestBed.createComponent(EditFinancialInstitutionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
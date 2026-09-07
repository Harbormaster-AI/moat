
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateDiagnosisComponent } from './create.component';
import { DiagnosisService } from '../../../services/Diagnosis.service';
import { Router } from '@angular/router';

describe('CreateDiagnosisComponent', () => {
  let component: CreateDiagnosisComponent;
  let fixture: ComponentFixture<CreateDiagnosisComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateDiagnosisComponent
      ],
      providers: [
        DiagnosisService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateDiagnosisComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
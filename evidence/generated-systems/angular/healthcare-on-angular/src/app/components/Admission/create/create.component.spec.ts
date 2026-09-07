
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateAdmissionComponent } from './create.component';
import { AdmissionService } from '../../../services/Admission.service';
import { Router } from '@angular/router';

describe('CreateAdmissionComponent', () => {
  let component: CreateAdmissionComponent;
  let fixture: ComponentFixture<CreateAdmissionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateAdmissionComponent
      ],
      providers: [
        AdmissionService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateAdmissionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
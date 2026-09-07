
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateClinicianComponent } from './create.component';
import { ClinicianService } from '../../../services/Clinician.service';
import { Router } from '@angular/router';

describe('CreateClinicianComponent', () => {
  let component: CreateClinicianComponent;
  let fixture: ComponentFixture<CreateClinicianComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateClinicianComponent
      ],
      providers: [
        ClinicianService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateClinicianComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
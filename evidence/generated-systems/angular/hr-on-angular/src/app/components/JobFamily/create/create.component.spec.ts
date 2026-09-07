
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateJobFamilyComponent } from './create.component';
import { JobFamilyService } from '../../../services/JobFamily.service';
import { Router } from '@angular/router';

describe('CreateJobFamilyComponent', () => {
  let component: CreateJobFamilyComponent;
  let fixture: ComponentFixture<CreateJobFamilyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateJobFamilyComponent
      ],
      providers: [
        JobFamilyService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateJobFamilyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
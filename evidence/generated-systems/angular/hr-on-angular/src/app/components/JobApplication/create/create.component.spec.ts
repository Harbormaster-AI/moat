
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateJobApplicationComponent } from './create.component';
import { JobApplicationService } from '../../../services/JobApplication.service';
import { Router } from '@angular/router';

describe('CreateJobApplicationComponent', () => {
  let component: CreateJobApplicationComponent;
  let fixture: ComponentFixture<CreateJobApplicationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateJobApplicationComponent
      ],
      providers: [
        JobApplicationService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateJobApplicationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
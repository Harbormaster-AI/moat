
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateOnboardingTaskComponent } from './create.component';
import { OnboardingTaskService } from '../../../services/OnboardingTask.service';
import { Router } from '@angular/router';

describe('CreateOnboardingTaskComponent', () => {
  let component: CreateOnboardingTaskComponent;
  let fixture: ComponentFixture<CreateOnboardingTaskComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateOnboardingTaskComponent
      ],
      providers: [
        OnboardingTaskService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateOnboardingTaskComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
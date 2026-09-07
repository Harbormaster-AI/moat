
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditOnboardingTaskComponent } from './edit.component';
import { OnboardingTaskService } from '../../../services/OnboardingTask.service';

describe('EditOnboardingTaskComponent', () => {
  let component: EditOnboardingTaskComponent;
  let fixture: ComponentFixture<EditOnboardingTaskComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditOnboardingTaskComponent
      ],
      providers: [
        OnboardingTaskService,
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

    fixture = TestBed.createComponent(EditOnboardingTaskComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
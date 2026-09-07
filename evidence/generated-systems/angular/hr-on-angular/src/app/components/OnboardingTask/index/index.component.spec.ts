
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexOnboardingTaskComponent } from './index.component';
import { OnboardingTaskService } from '../../../services/OnboardingTask.service';

describe('IndexOnboardingTaskComponent', () => {
  let component: IndexOnboardingTaskComponent;
  let fixture: ComponentFixture<IndexOnboardingTaskComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexOnboardingTaskComponent
      ],
      providers: [
        OnboardingTaskService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexOnboardingTaskComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
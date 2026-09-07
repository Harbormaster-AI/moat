
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateTrainingRunComponent } from './create.component';
import { TrainingRunService } from '../../../services/TrainingRun.service';
import { Router } from '@angular/router';

describe('CreateTrainingRunComponent', () => {
  let component: CreateTrainingRunComponent;
  let fixture: ComponentFixture<CreateTrainingRunComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateTrainingRunComponent
      ],
      providers: [
        TrainingRunService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateTrainingRunComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
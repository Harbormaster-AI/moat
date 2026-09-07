
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexTrainingRunComponent } from './index.component';
import { TrainingRunService } from '../../../services/TrainingRun.service';

describe('IndexTrainingRunComponent', () => {
  let component: IndexTrainingRunComponent;
  let fixture: ComponentFixture<IndexTrainingRunComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexTrainingRunComponent
      ],
      providers: [
        TrainingRunService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexTrainingRunComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
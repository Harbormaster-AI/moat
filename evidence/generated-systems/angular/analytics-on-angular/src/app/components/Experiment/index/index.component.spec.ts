
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexExperimentComponent } from './index.component';
import { ExperimentService } from '../../../services/Experiment.service';

describe('IndexExperimentComponent', () => {
  let component: IndexExperimentComponent;
  let fixture: ComponentFixture<IndexExperimentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexExperimentComponent
      ],
      providers: [
        ExperimentService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexExperimentComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexRecommendationScenarioComponent } from './index.component';
import { RecommendationScenarioService } from '../../../services/RecommendationScenario.service';

describe('IndexRecommendationScenarioComponent', () => {
  let component: IndexRecommendationScenarioComponent;
  let fixture: ComponentFixture<IndexRecommendationScenarioComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexRecommendationScenarioComponent
      ],
      providers: [
        RecommendationScenarioService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexRecommendationScenarioComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
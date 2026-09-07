
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexVisualizationComponent } from './index.component';
import { VisualizationService } from '../../../services/Visualization.service';

describe('IndexVisualizationComponent', () => {
  let component: IndexVisualizationComponent;
  let fixture: ComponentFixture<IndexVisualizationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexVisualizationComponent
      ],
      providers: [
        VisualizationService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexVisualizationComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
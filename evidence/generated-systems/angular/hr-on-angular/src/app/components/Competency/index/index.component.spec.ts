
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexCompetencyComponent } from './index.component';
import { CompetencyService } from '../../../services/Competency.service';

describe('IndexCompetencyComponent', () => {
  let component: IndexCompetencyComponent;
  let fixture: ComponentFixture<IndexCompetencyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexCompetencyComponent
      ],
      providers: [
        CompetencyService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexCompetencyComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
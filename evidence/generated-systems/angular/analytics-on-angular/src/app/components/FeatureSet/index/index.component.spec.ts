
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexFeatureSetComponent } from './index.component';
import { FeatureSetService } from '../../../services/FeatureSet.service';

describe('IndexFeatureSetComponent', () => {
  let component: IndexFeatureSetComponent;
  let fixture: ComponentFixture<IndexFeatureSetComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexFeatureSetComponent
      ],
      providers: [
        FeatureSetService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexFeatureSetComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
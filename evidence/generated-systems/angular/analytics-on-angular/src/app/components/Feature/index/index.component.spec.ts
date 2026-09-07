
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexFeatureComponent } from './index.component';
import { FeatureService } from '../../../services/Feature.service';

describe('IndexFeatureComponent', () => {
  let component: IndexFeatureComponent;
  let fixture: ComponentFixture<IndexFeatureComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexFeatureComponent
      ],
      providers: [
        FeatureService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexFeatureComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
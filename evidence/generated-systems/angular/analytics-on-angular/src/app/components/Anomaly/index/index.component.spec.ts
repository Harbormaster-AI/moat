
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexAnomalyComponent } from './index.component';
import { AnomalyService } from '../../../services/Anomaly.service';

describe('IndexAnomalyComponent', () => {
  let component: IndexAnomalyComponent;
  let fixture: ComponentFixture<IndexAnomalyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexAnomalyComponent
      ],
      providers: [
        AnomalyService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexAnomalyComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
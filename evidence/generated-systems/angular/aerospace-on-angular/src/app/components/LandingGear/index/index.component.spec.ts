
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexLandingGearComponent } from './index.component';
import { LandingGearService } from '../../../services/LandingGear.service';

describe('IndexLandingGearComponent', () => {
  let component: IndexLandingGearComponent;
  let fixture: ComponentFixture<IndexLandingGearComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexLandingGearComponent
      ],
      providers: [
        LandingGearService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexLandingGearComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
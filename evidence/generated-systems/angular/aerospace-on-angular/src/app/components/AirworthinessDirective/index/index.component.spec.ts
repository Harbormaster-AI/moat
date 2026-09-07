
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexAirworthinessDirectiveComponent } from './index.component';
import { AirworthinessDirectiveService } from '../../../services/AirworthinessDirective.service';

describe('IndexAirworthinessDirectiveComponent', () => {
  let component: IndexAirworthinessDirectiveComponent;
  let fixture: ComponentFixture<IndexAirworthinessDirectiveComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexAirworthinessDirectiveComponent
      ],
      providers: [
        AirworthinessDirectiveService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexAirworthinessDirectiveComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
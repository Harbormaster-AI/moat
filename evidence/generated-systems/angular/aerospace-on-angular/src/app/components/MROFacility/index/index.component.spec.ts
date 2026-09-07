
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexMROFacilityComponent } from './index.component';
import { MROFacilityService } from '../../../services/MROFacility.service';

describe('IndexMROFacilityComponent', () => {
  let component: IndexMROFacilityComponent;
  let fixture: ComponentFixture<IndexMROFacilityComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexMROFacilityComponent
      ],
      providers: [
        MROFacilityService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexMROFacilityComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexCabinLayoutComponent } from './index.component';
import { CabinLayoutService } from '../../../services/CabinLayout.service';

describe('IndexCabinLayoutComponent', () => {
  let component: IndexCabinLayoutComponent;
  let fixture: ComponentFixture<IndexCabinLayoutComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexCabinLayoutComponent
      ],
      providers: [
        CabinLayoutService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexCabinLayoutComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
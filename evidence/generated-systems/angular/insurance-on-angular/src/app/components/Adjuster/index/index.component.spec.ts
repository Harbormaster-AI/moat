
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexAdjusterComponent } from './index.component';
import { AdjusterService } from '../../../services/Adjuster.service';

describe('IndexAdjusterComponent', () => {
  let component: IndexAdjusterComponent;
  let fixture: ComponentFixture<IndexAdjusterComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexAdjusterComponent
      ],
      providers: [
        AdjusterService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexAdjusterComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});